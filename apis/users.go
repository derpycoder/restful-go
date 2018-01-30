package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var projectID = "chrome-setup-158308"

type User struct {
	// Auto Generated
	// Read Only: true
	ID *datastore.Key `datastore:"__key__" json:"id"`

	// From where to inherit
	// Required: true
	ParentID string `datastore:"-" json:"parent_id"`

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
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	b, _ := ioutil.ReadAll(r.Body)

	var user = new(User)

	err = json.Unmarshal(b, user)
	if err != nil {
		http.Error(w, "Unable to parse Request BODY: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	parentKey, err := datastore.DecodeKey(user.ParentID)
	if err != nil {
		http.Error(w, "Unable to Decode Parent ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	key := datastore.IDKey("Users", 0, parentKey)
	key.Namespace = "NeverLand"

	key, err = client.Put(ctx, key, user)
	if err != nil {
		http.Error(w, "Failed to create User: "+err.Error(), http.StatusExpectationFailed)
		return
	}

	user.ID = key

	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	query := datastore.NewQuery("Users").Limit(10).Namespace("NeverLand")

	var users []*User

	keys, err := client.GetAll(ctx, query, &users)
	if err != nil {
		http.Error(w, "No Users Found: "+err.Error(), http.StatusNotFound)
		return
	}

	json, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(keys) > 0 {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, "Failed to establish connection with Database: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	var user = new(User)

	key, err := datastore.DecodeKey(params["id"])
	if err != nil {
		http.Error(w, "Unable to Decode ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = client.Get(ctx, key, user)
	if err != nil {
		http.Error(w, "User Not Found: "+err.Error(), http.StatusNotFound)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Unable to convert response to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
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
