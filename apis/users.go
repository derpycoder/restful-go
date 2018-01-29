package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

type User struct {
	// ID can be used for further queries
	// Read Only: true
	UUID string `datastore:"uuid" json:"uuid"`

	// First Name of the User
	// Required: true
	Firstname string `datastore:"firstname" json:"firstname"`

	// Middle Name of the User
	Middlename string `datastore:"middlename,omitempty" json:"middlename,omitempty"`

	// Last Name of the User
	Lastname string `datastore:"lastname,omitempty" json:"lastname,omitempty"`

	// Validated Email Ids
	EmailIds []string `datastore:"email_ids" json:"email_ids"`

	// 10 Digit Unformatted Phone Numbers
	PhoneNumbers []string `datastore:"phone_numbers" json:"phone_numbers"`

	// Date of Birth in UTC Zulu Format
	// Required: true
	Dob time.Time `datastore:"dob" json:"dob"`

	// Role of the User
	Role string `datastore:"role,omitempty" json:"role,omitempty"`

	// User's Profile Photo, Needs to be Uploaded
	// Read Only: true
	AvatarURL string `datastore:"avatar_url,omitempty" json:"avatar_url,omitempty"`

	// Blood Group of the User
	BloodGroup string `datastore:"blood_group,omitempty" json:"blood_group,omitempty"`

	// UUID of the Society the User belongs to
	// Required: true
	SocietyUUID string `datastore:"society_uuid" json:"society_uuid"`

	// UUID of the Family the User belongs to
	// Required: true
	FamilyUUID string `datastore:"family_uuid" json:"family_uuid"`

	// Date in UTC Zulu Format
	// Read Only: true
	DateCreated time.Time `datastore:"date_created" json:"date_created"`

	// Date in UTC Zulu Format
	// Read Only: true
	DateUpdated time.Time `datastore:"date_updated" json:"date_updated"`

	// Status of the User's Account
	// Read Only: true
	Status string `datastore:"status,omitempty" json:"status,omitempty"`
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

var projectID = "chrome-setup-158308"

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Set your Google Cloud Platform project ID.

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
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		fmt.Printf("Failed")
	}

	var user = new(User)

	key := datastore.NameKey("Users", params["uuid"], nil)

	key.Namespace = "NeverLand"

	err = client.Get(ctx, key, user)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		log.Fatalf("Unable to Get: %v", err)
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Bad Result: %v", err)
		w.WriteHeader(http.StatusGone)
	}
	w.Write(json)
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
