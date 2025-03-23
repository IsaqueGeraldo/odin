package odin

import (
	"log"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Envoriment struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var conn *gorm.DB

func isValidFileName(fileName string) bool {
	if fileName == "" {
		return false
	}

	invalidChars := `[\\/:*?"<>|]`
	matched, _ := regexp.MatchString(invalidChars, fileName)
	if matched {
		return false
	}

	if len(fileName) > 255 {
		return false
	}

	if strings.HasSuffix(fileName, " ") || strings.HasSuffix(fileName, ".") {
		return false
	}

	return true
}

func Bootstrap(source ...string) {
	var err error

	dsn := "odin.db"
	if len(source) > 0 {
		if isValidFileName(source[0]) {
			dsn = source[0]
		}
	}

	conn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Printf("failed to retrieve SQL instance from GORM: %v", err)
		return
	}

	if err = sqlDB.Ping(); err != nil {
		log.Printf("failed to ping database: %v", err)
		return
	}

	conn.AutoMigrate(&Envoriment{})
}

func sanitizeKey(key string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, key)
	if err != nil {
		return key
	}

	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	sanitized := re.ReplaceAllString(result, "_")

	return strings.ToUpper(sanitized)
}

func Getenv(key string) string {
	key = sanitizeKey(key)
	var env Envoriment
	if err := conn.Where("key = ?", key).First(&env).Error; err != nil {
		return ""
	}
	return env.Value
}

func Setenv(key string, value string) error {
	key = sanitizeKey(key)
	env := Envoriment{Key: key, Value: value}

	err := conn.Where("key = ?", key).First(&env).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return conn.Create(&env).Error
		}
		return err
	}

	return conn.Where("key = ?", key).Save(&env).Error
}

func Unsetenv(key string) error {
	key = sanitizeKey(key)
	return conn.Where("key = ?", key).Delete(&Envoriment{}).Error
}

func Clearenv() {
	conn.Exec("DELETE FROM environments")
}

func Environ() []string {
	var envs []Envoriment
	if err := conn.Find(&envs).Error; err != nil {
		log.Printf("failed to retrieve environment variables: %v", err)
		return nil
	}

	var result []string
	for _, env := range envs {
		result = append(result, env.Key+"="+env.Value)
	}
	return result
}

func ExpandEnv(str string) string {
	envs := Environ()
	for _, env := range envs {
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]
		str = strings.ReplaceAll(str, "$"+key, value)
		str = strings.ReplaceAll(str, "${"+key+"}", value)
	}
	return str
}

func LookupEnv(key string) (string, bool) {
	key = sanitizeKey(key)
	value := Getenv(key)
	return value, value != ""
}
