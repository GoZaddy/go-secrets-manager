package vault

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	pathLib "path"

	"github.com/gozaddy/secret.ly/mycrypto"
	homedir "github.com/mitchellh/go-homedir"
)

/*Vault represents a vault where api keys and similar sensitive information can be stored and retrieved.
To open a vault, use FileVault
*/
type Vault struct {
	data     map[string]string
	FilePath string
}

//FileVaultOptions defines options for the FileVault function
type FileVaultOptions struct {
	CreateNew bool
}

/*
FileVault takes in an encoding key(similar to a password) and file path to the vault. If both parameters are correct and valid, a vault is returned.
*/
func FileVault(filePath string, options FileVaultOptions) (Vault, error) {
	var dataMap map[string]string
	var data []byte

	//get home directory
	home, err := homedir.Dir()
	if err != nil {
		return Vault{}, err
	}

	//make secret.ly directory
	path := pathLib.Join(home, "secret.ly")
	err = os.MkdirAll(path, 0660) //creates path if it doesn't exist
	if err != nil {
		return Vault{}, err
	}

	//read/create file
	path = pathLib.Join(path, filePath)

	for {
		data, err = ioutil.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) && options.CreateNew == true {
				_, err = os.Create(path)
				if err != nil {
					return Vault{}, err
				}
				err = writeEmptyMapToFile(path)
				if err != nil {
					return Vault{}, err
				}
			} else {
				return Vault{}, err
			}

		} else {
			break
		}
	}

	//write empty map to file if file's empty
	if string(data) == "" {
		err = writeEmptyMapToFile(path)
		if err != nil {
			return Vault{}, err
		}
		data, err = ioutil.ReadFile(path)
		if err != nil {
			return Vault{}, err
		}
	}

	err = json.Unmarshal(data, &dataMap)
	if err != nil {

		return Vault{}, err
	}

	return Vault{data: dataMap, FilePath: path}, nil

}

//Get retrives a secret using a keyname
func (v Vault) Get(keyname, encodingKey string) (string, error) {
	result, err := mycrypto.Decrypt(encodingKey, hex.EncodeToString([]byte(v.data[keyname])))
	return result, err
}

//Set Stores a secret using the keyname as an identifier
func (v Vault) Set(keyname, keyvalue, encodingKey string) error {
	result, err := mycrypto.Encrypt(encodingKey, keyvalue)
	if err != nil {
		return err
	}
	v.data[keyname] = result
	bs, err := json.Marshal(v.data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(v.FilePath, bs, 0660)
	if err != nil {
		return err
	}
	return nil

}

func writeEmptyMapToFile(path string) error {
	bs, err := json.Marshal(map[string]string{})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bs, 0660)
	if err != nil {
		return err
	}
	return nil
}
