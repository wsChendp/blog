package flag

import (
	"os"
	"server/global"
	"strings"
)

// SQLImport ���� MySQL ����
func SQLImport(sqlPath string) (errs []error) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		return append(errs, err)
	}
	// �ָ�����
	sqlList := strings.Split(string(byteData), ";")
	for _, sql := range sqlList {
		// ȥ���ַ�����ͷ�ͽ�β�Ŀհ׷�
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		// ִ��sql���
		err = global.DB.Exec(sql).Error
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return nil
}