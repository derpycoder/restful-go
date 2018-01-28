package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"google.golang.org/appengine"
)

type User struct {
	AvatarURL string `datastore:"avatar_url,omitempty" json:"avatar_url,omitempty"`

	// Blood Group of the User
	BloodGroup string `datastore:"blood_group,omitempty" json:"blood_group,omitempty`

	// Date in UTC Zulu Format
	// Required: true
	DateCreated time.Time `datastore:"date_created" json:"date_created"`

	// Date in UTC Zulu Format
	// Required: true
	DateUpdated time.Time `datastore:"date_updated" json:"date_updated"`

	// Date of Birth in UTC Zulu Format
	// Required: true
	Dob time.Time `datastore:"dob" json:"dob"`

	// Validated Email Ids
	EmailIds []string `datastore:"email_ids" json:"email_ids"`

	// ID of the Family the User belongs to
	// Required: true
	FamilyID string `datastore:"family_id" json:"family_id"`

	// First Name of the User
	// Required: true
	Firstname string `datastore:"firstname" json:"firstname"`

	// Last Name of the User
	Lastname string `datastore:"lastname,omitempty" json:"lastname,omitempty"`

	// Middle Name of the User
	Middlename string `datastore:"middlename,omitempty" json:"middlename,omitempty"`

	// 10 Digit Unformatted Phone Numbers
	PhoneNumbers []string `datastore:"phone_numbers" json:"phone_numbers"`

	// Role of the User
	Role string `datastore:"role,omitempty" json:"role,omitempty"`

	// ID of the Society the User belongs to
	// Required: true
	SocietyID string `datastore:"society_id" json:"society_id"`

	// Status of the User's Account
	Status string `datastore:"status,omitempty" json:"status,omitempty"`

	// ID can be used for further queries
	// Read Only: true
	id *datastore.Key `datastore:"__key__" json:"id"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type marshalledUser User
	return json.Marshal(&struct {
		Dob string `datastore:"dob" json:"dob"`
		*marshalledUser
	}{
		Dob:            time.Time(u.Dob).Format("2006-01-02"),
		marshalledUser: (*marshalledUser)(u),
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Set your Google Cloud Platform project ID.
	projectID := "chrome-setup-158308"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		fmt.Printf("Failed")
	}

	query := datastore.NewQuery("Users").Limit(10).Namespace("NeverLand")

	var users []*User

	keys, err := client.GetAll(ctx, query, &users)
	if err != nil {
		w.Write([]byte("Errrrrr: " + err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if len(keys) > 0 {
		w.Write(json)
		w.WriteHeader(http.StatusOK)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
