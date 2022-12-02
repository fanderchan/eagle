package db

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"

	"eagle/util"
)

const defaultFileName = "eagle.json"

type Cache struct {
	filePath string
}

var c *Cache

func GetDb() (*Cache, error) {
	if c != nil {
		return c, nil
	}

	c = &Cache{}
	filePath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	c.filePath = filePath

	return c, nil

}

func getFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := path.Join(dir, defaultFileName)
	if util.CheckFileIsExist(filePath) {
		return filePath, nil
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", nil
	}

	defer file.Close()
	return filePath, nil

}

func (c *Cache) rewrite(body []byte) error {
	file, err := os.OpenFile(c.filePath, os.O_RDWR, 0644)
	defer file.Close()
	file.Seek(0, 0)
	file.Truncate(0)
	_, err = file.Write(body)
	return err
}

func (c *Cache) getFileData() (map[string]interface{}, error) {
	file, err := os.OpenFile(c.filePath, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *Cache) Save(key string, data interface{}) error {

	var (
		file *os.File
		err  error
	)

	file, err = os.OpenFile(c.filePath, os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	fb, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fb) == 0 {
		// 文件为空直接写入
		m := make(map[string]interface{})
		m[key] = data
		body, err := json.Marshal(m)
		if err != nil {
			return err
		}
		_, err = file.Write(body)
		return err
	}

	// 文件不为空,先解析
	md := make(map[string]interface{})
	err = json.Unmarshal(fb, &md)
	if err != nil {
		return err
	}
	_, ok := md[key]
	if ok {
		return errors.New("duplicate record")
	}

	md[key] = data
	body, err := json.Marshal(md)
	// 写之前位移到文件开头，并且将之前的文件内容清楚
	file.Seek(0, 0)
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	return err
}

func (c *Cache) Get(key string) (interface{}, error) {
	m, err := c.getFileData()
	if err != nil {
		return nil, err
	}
	value, ok := m[key]

	if !ok {
		return nil, errors.New("record not found")
	}

	return value, nil
}

func (c *Cache) Delete(key string) error {
	m, err := c.getFileData()
	if err != nil {
		return err
	}

	_, ok := m[key]
	if !ok {
		return errors.New("record not found")
	}

	delete(m, key)
	// 写入文件
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = c.rewrite(body)

	return err
}

func (c *Cache) Update(key string, data interface{}) error {
	m, err := c.getFileData()
	if err != nil {
		return err
	}
	m[key] = data
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return c.rewrite(body)
}

func (c *Cache) List() (interface{}, error) {
	m, err := c.getFileData()
	if err != nil {
		return nil, err
	}

	return m, err
}
