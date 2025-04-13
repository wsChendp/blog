package flag

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"time"

	"go.uber.org/zap"
)

// SQLExport 导出 MySQL 数据
func SQLExport() error {
	mysql := global.Config.Mysql

	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("mysql_%s.sql", timer)

	// 提示用户输入服务器密码
	global.Log.Info("将使用PowerShell连接到阿里云服务器进行备份")
	fmt.Print("请输入阿里云服务器密码: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sshPassword := scanner.Text()

	// 生成PowerShell远程连接命令，在阿里云服务器上执行备份
	global.Log.Info("正在连接服务器并执行备份...")

	// 创建PowerShell命令
	// 1. 首先在服务器上执行mysqldump
	// 2. 然后将文件下载到本地
	psCommand := fmt.Sprintf(`
$SecurePassword = ConvertTo-SecureString "%s" -AsPlainText -Force
$Credential = New-Object System.Management.Automation.PSCredential ("root", $SecurePassword)
$Session = New-PSSession -ComputerName %s -Credential $Credential

Invoke-Command -Session $Session -ScriptBlock {
    mysqldump -u%s -p%s --no-tablespaces %s > /tmp/backup.sql
    Write-Host "备份已完成，文件保存在服务器/tmp/backup.sql"
}

Remove-PSSession $Session

# 下载文件
$SecurePassword = ConvertTo-SecureString "%s" -AsPlainText -Force
$Credential = New-Object System.Management.Automation.PSCredential ("root", $SecurePassword)
$Session = New-PSSession -ComputerName %s -Credential $Credential

Copy-Item -Path "/tmp/backup.sql" -Destination "./%s" -FromSession $Session
Remove-PSSession $Session
Write-Host "文件已下载到本地"
`, sshPassword, mysql.Host, mysql.Username, mysql.Password, mysql.DBName, sshPassword, mysql.Host, sqlPath)

	// 创建临时PS脚本文件
	tempPsFile := "temp_backup.ps1"
	err := os.WriteFile(tempPsFile, []byte(psCommand), 0644)
	if err != nil {
		global.Log.Error("创建PS脚本失败", zap.Error(err))
		return err
	}
	defer os.Remove(tempPsFile)

	// 执行PowerShell脚本
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", tempPsFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		global.Log.Error("执行PowerShell命令失败", zap.Error(err))
		global.Log.Info("PowerShell远程连接可能需要额外配置，请尝试手动备份:")
		global.Log.Info("1. 使用SSH登录到服务器: ssh root@" + mysql.Host)
		global.Log.Info(fmt.Sprintf("2. 执行备份命令: mysqldump -u%s -p%s --no-tablespaces %s > /tmp/backup.sql",
			mysql.Username, mysql.Password, mysql.DBName))
		global.Log.Info("3. 下载备份文件: scp root@" + mysql.Host + ":/tmp/backup.sql ./" + sqlPath)
		return err
	}

	global.Log.Info("备份成功完成！SQL文件已保存到: " + sqlPath)
	return nil
}
