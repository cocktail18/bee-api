package util

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"os/exec"
	"time"
)

func StopDockerMysqlForTest(containerName string) error {
	cmd := exec.Command("docker", "stop", containerName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "Failed to stop MySQL container")
	}
	time.Sleep(time.Second * 2)
	cmd = exec.Command("docker", "rm", containerName)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "Failed to remove MySQL container")
	}
	time.Sleep(time.Second * 2)
	return nil
}

func InitDockerMysqlForTest() (string, string, error) {
	port, err := GetFreePort()
	if err != nil {
		return "", "", err
	}
	// 指定要运行的 MySQL 容器的名称
	containerName := fmt.Sprintf("mysql-for-unit-test-" + cast.ToString(port))
	pass := "123456"
	// 执行 docker run 命令启动 MySQL 容器
	cmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%d:3306", port), "--name", containerName, "-e", "MYSQL_ROOT_PASSWORD="+pass, "mysql:8.4.1")

	// 执行命令并捕获输出和错误
	_, err = cmd.CombinedOutput()
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to start MySQL container")
	}
	fmt.Println("等待MySQL启动")
	<-time.After(time.Second * 10)
	// 连接 MySQL 数据库
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(localhost:%d)/", pass, port))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建数据库
	_, err = db.Exec("CREATE DATABASE bee")
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to create database")
	}
	return containerName, fmt.Sprintf("root:%s@tcp(localhost:%d)/bee?charset=utf8mb4&parseTime=True&loc=Local", pass, port), nil
}
