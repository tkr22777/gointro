package simplepackage

import (
	"context"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stupefied/simplepackage/ent"
	"github.com/stupefied/simplepackage/ent/user"
)

//followed https://entgo.io/docs/getting-started/

func TestEnt(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	createdUser, err := CreateUser(context.TODO(), client, &ent.User{Name: "test-user", Age: 76})
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	fmt.Println("User created:", createdUser.String())

	retrivedUser, err := GetUser(context.TODO(), client, createdUser.ID)
	if err != nil {
		log.Fatalf("Could not retrieve user: %v", err)
	}

	fmt.Println("User retrieved:", retrivedUser.String())
}

func TestEntSearch(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	data := [][]string{
		{"John", "innovate@gmail.com", ""},
		{"Tahsin", "tkr@gmail.com", ""},
		{"Rohn", "sans@hotmail.com", ""},
		{"Mohsin", "mohn@yahoo.com", ""},
		{"Ahsan", "dohn@hotmail.com", ""},
	}

	for i := 0; i < len(data); i++ {
		createdUser, err := CreateUser(context.TODO(), client,
			&ent.User{
				Age:     i + 1,
				Name:    data[i][0],
				Email:   data[i][1],
				Address: data[i][2],
			},
		)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
		fmt.Println("User created:", createdUser.String())
	}

	users, _ := SearchUserByName(context.TODO(), client, "in")
	for i := 0; i < len(users); i++ {
		fmt.Println("User found:", users[i].String())
	}
}

func CreateUser(ctx context.Context, client *ent.Client, u *ent.User) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(u.Age).
		SetName(u.Name).
		SetEmail(u.Email).
		SetAddress(u.Address).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}

func GetUser(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	user, err := client.User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive user: %w", err)
	}
	return user, nil
}

func SearchUserByName(ctx context.Context, client *ent.Client, searchTerm string) ([]*ent.User, error) {
	users, err := client.User.Query().Where(user.NameContains(searchTerm)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive user: %w", err)
	}
	return users, nil
}

func SearchUserByTemp(t *testing.T) {
	fmt.Printf(user.AddressHasPrefix("blaj").String())
}

func SearchUserByNameOrEmail(ctx context.Context, client *ent.Client, searchTerm string) ([]*ent.User, error) {
	users, err := client.User.Query().
		Where(
			predicate.User(func(s *sql.Selector) {
				s.Where(
				),
			}),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive user: %w", err)
	}
	return users, nil
}
