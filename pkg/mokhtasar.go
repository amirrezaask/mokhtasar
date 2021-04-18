package pkg

import "database/sql"

type Mokhtasar struct {
	DB              *sql.DB //TODO: abstract repository layer
	RandomGenerator func(len int) string
}

func (m *Mokhtasar) GetOriginalURL(key string) (string, error) {
	query := `SELECT url FROM urls WHERE key=$1`
	rows, err := m.DB.Query(query, key)
	if err != nil {
		return "", err
	}
	var url string
	for rows.Next() {
		err = rows.Scan(&url)
		if err != nil {
			return "", err
		}
	}
	return url, nil

}
func (m *Mokhtasar) Shorten(url string) (string, error) {
	key := m.RandomGenerator(5)
	_, err := m.DB.Exec(`INSERT INTO urls (url, key) VALUES ($1, $2)`, url, key)
	if err != nil {
		return "", err
	}
	return key, nil

}
