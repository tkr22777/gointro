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

	createdUser, err := CreateUser(context.TODO(), client, "test-user", 76)
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

	names := []string{
		"John",
		"Tahsin",
		"Ringo",
		"Paul",
		"Anita",
		"George",
		"Chiesa",
		"Gregor",
		"Manchini",
		"Stephenie",
	}

	for i := 0; i < len(names); i++ {
		createdUser, err := CreateUser(context.TODO(), client, names[i], i+1)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
		fmt.Println("User created:", createdUser.String())
	}

	users, _ := SearchUser(context.TODO(), client, "in")
	for i := 0; i < len(users); i++ {
		fmt.Println("User found:", users[i].String())
	}

}

func CreateUser(ctx context.Context, client *ent.Client, name string, age int) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(age).
		SetName(name).
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

func SearchUser(ctx context.Context, client *ent.Client, searchTerm string) ([]*ent.User, error) {
	users, err := client.User.Query().Where(user.NameContains(searchTerm)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive user: %w", err)
	}
	return users, nil
}
