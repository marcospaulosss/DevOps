package repositories

import (
	"fmt"
)

type AccountQueryBuilder struct{}

func NewAccountQueryBuilder() AccountQueryBuilder {
	return AccountQueryBuilder{}
}

func (this AccountQueryBuilder) CreateAccount(bind string) string {
	return fmt.Sprintf(`INSERT INTO accounts (%[1]s, %[1]s_code)
		SELECT :%[1]s, :%[1]s_code
		WHERE NOT EXISTS (
			SELECT id FROM users WHERE %[1]s ilike :%[1]s
		) RETURNING id;`, bind)
}

func (this AccountQueryBuilder) ValidateAccount(bind string) string {
	return fmt.Sprintf(`SELECT p.*, COALESCE((select true from users as u where u.%[1]s = :%[1]s), false) as exists,
				     coalesce(
           				cast(
					    (select id from users as u where u.%[1]s = :%[1]s) as text
               				), ''
				     ) as user_id
 				     FROM accounts as p
				     WHERE p.%[1]s = :%[1]s
				     AND p.%[1]s_code = :%[1]s_code limit 1;`, bind)
}

func (this AccountQueryBuilder) DeleteAccount() string {
	return "DELETE FROM accounts WHERE (phone = :phone AND phone != '') OR (email = :email AND email != '') RETURNING id"
}

func (this AccountQueryBuilder) SelectAccountByType(bind string) string {
	return fmt.Sprintf("SELECT *, COALESCE((select true from users as u where u.%[1]s = :%[1]s), false) as exists FROM accounts WHERE %[1]s ilike :%[1]s", bind)
}

func (this AccountQueryBuilder) CreateAccountLogin(bind string) string {
	return fmt.Sprintf(`INSERT INTO accounts (%[1]s, %[1]s_code) VALUES
		(:%[1]s, :%[1]s_code) RETURNING id`, bind)
}
