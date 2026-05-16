package bw

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Client struct {
	env           map[string]string
	sessionRegexp *regexp.Regexp
}

func NewClient(serverUrl string, email string, password string) (*client, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	appDataDirName := strings.ReplaceAll(serverUrl, "https://", "")
	appDataDirName = strings.ReplaceAll(appDataDirName, "http://", "")
	appDataDirName = strings.ReplaceAll(appDataDirName, ".", "-")
	appDataDirName = strings.ReplaceAll(appDataDirName, "/", "-")

	appDataDir := fmt.Sprintf("%s/.kuhweide/go-bw-lib/%s", homeDir, appDataDirName)

	client := &client{}

	client.env = map[string]string{"BITWARDENCLI_APPDATA_DIR": appDataDir}
	client.sessionRegexp = regexp.MustCompile(`BW_SESSION="(\S+?)"`)

	clientStatus, err := client.Status()
	if err != nil {
		return nil, fmt.Errorf("Failed to get status for initializing client. err=%w", err)
	}

	if clientStatus.UserEmail != nil && *clientStatus.UserEmail != email {
		err := client.Logout()

		if err != nil {
			err := fmt.Errorf("Failed to change email, because logout failed. err=%w", err)
			println(err)
			return nil, err
		}
	}

	clientStatus, err = client.Status()
	if err != nil {
		return nil, fmt.Errorf("Failed to get status for initializing client. err=%w", err)
	}

	if clientStatus.ServerUrl != serverUrl {
		if clientStatus.Status == LockStatusLocked || clientStatus.Status == LockStatusUnlocked {
			err := client.Logout()

			if err != nil {
				err := fmt.Errorf("Failed to set server, because logout failed. err=%w", err)
				println(err)
				return nil, err
			}
		}

		client.SetServerUrl(serverUrl)
	}

	if clientStatus.Status == LockStatusUnauthenticated {
		err = client.Login(email, password)
		if err != nil {
			return nil, err
		}
	} else {
		err = client.Unlock(password)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

func (client *client) generateEnv() []string {
	env := make([]string, 0)

	for key, value := range client.env {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	return env
}

func (client *client) command(arg ...string) (string, error) {
	cmd := exec.Command("bw", arg...)
	cmd.Env = client.generateEnv()

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Failed to run command. out=%s, err=%s", out, err)
	}

	return strings.TrimSpace(string(out[:])), nil
}

func (client *client) commandWithPipedInput(input string, arg ...string) (string, error) {
	cmd := exec.Command("bw", arg...)
	cmd.Env = client.generateEnv()
	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Failed to run command. out=%s, err=%s", out, err)
	}

	return strings.TrimSpace(string(out[:])), nil
}

func (client *client) Login(email string, password string) error {
	out, err := client.command("login", email, password)

	if err != nil {
		err = fmt.Errorf("Failed to login. err=%w", err)
		fmt.Println(err)
		return err
	}

	client.env["BW_SESSION"] = client.sessionRegexp.FindStringSubmatch(out)[1]

	return nil
}

func (client *client) Unlock(password string) error {
	out, err := client.command("unlock", password)

	if err != nil {
		err = fmt.Errorf("Failed to unlock. err=%w", err)
		fmt.Println(err)
		return err
	}

	client.env["BW_SESSION"] = client.sessionRegexp.FindStringSubmatch(out)[1]

	return nil
}

func (client *client) Logout() error {
	_, err := client.command("logout")

	if err != nil {
		err = fmt.Errorf("Failed to logout. err=%w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

func (client *client) SetServerUrl(server string) error {
	_, err := client.command("config", "server", server)

	if err != nil {
		err = fmt.Errorf("Failed to set server. err=%w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

func (client *client) Status() (*clientStatus, error) {
	out, err := client.command("status")

	if err != nil {
		err := fmt.Errorf("Failed to get client status. err=%w", err)
		fmt.Println(err.Error())

		return nil, err
	}

	lines := strings.Split(out, "\n")

	clientStatus := &clientStatus{}
	err = json.Unmarshal([]byte(strings.TrimSpace(lines[len(lines)-1])), clientStatus)

	if err != nil {
		err := fmt.Errorf("Failed to parse client status from response. err=%w", err)
		fmt.Println(err)
		return nil, err
	}

	return clientStatus, nil
}

func (client *client) Sync() error {
	_, err := client.command("sync")

	if err != nil {
		err := fmt.Errorf("Failed to get sync vault. err=%w", err)
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (client *client) Encode(s string) (string, error) {
	out, err := client.commandWithPipedInput(s, "encode")

	if err != nil {
		err := fmt.Errorf("Failed to encode string. err=%w", err)
		fmt.Println(err)
		return "", err
	}

	return out, nil
}

func (client *client) CreateItem(newItem *item) (*item, error) {
	data, err := json.Marshal(newItem)
	if err != nil {
		err := fmt.Errorf("Failed create item because generating json failed. err=%w", err)
		fmt.Println(err)
		return nil, err
	}

	encodedJson, err := client.Encode(string(data[:]))
	if err != nil {
		err := fmt.Errorf("Failed create item, because encoding json failed. err=%w", err)
		fmt.Println(err)
		return nil, err
	}

	out, err := client.command("create", "item", encodedJson)
	if err != nil {
		err := fmt.Errorf("Failed create item. err=%w", err)
		fmt.Println(err)
		return nil, err
	}

	createdItem := &item{}
	err = json.Unmarshal([]byte(out), createdItem)
	if err != nil {
		err := fmt.Errorf("Failed parse item from response. err=%w", err)
		fmt.Println(err)
		return nil, err
	}

	return createdItem, nil
}

func (client *client) Item(id string) (*item, error) {
	item := &item{}

	out, err := client.command("get", "item", id)

	if err != nil {
		err := fmt.Errorf("Failed to get sync vault. err=%w", err)
		fmt.Println(err.Error())
		return nil, err
	}

	err = json.Unmarshal([]byte(out), &item)

	return item, nil
}

func (client *client) Items() ([]item, error) {
	var items []item

	out, err := client.command("list", "items")

	if err != nil {
		err := fmt.Errorf("Failed to get sync vault. err=%w", err)
		fmt.Println(err.Error())
		return items, err
	}

	err = json.Unmarshal([]byte(out), &items)
	if err != nil {
		err := fmt.Errorf("Failed parse items from response. err=%w", err)
		fmt.Println(err)
		return nil, err
	}

	return items, nil
}
