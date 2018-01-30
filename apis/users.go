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
		w.Write([]byte("Failed to create client: " + err.Error()))
		appendHeader(w, http.StatusServiceUnavailable)
		return
	}

	b, _ := ioutil.ReadAll(r.Body)

	var user = new(User)

	err = json.Unmarshal(b, user)
	if err != nil {
		w.Write([]byte("Unable to Unmarshal: " + err.Error()))
		appendHeader(w, http.StatusServiceUnavailable)
		return
	}

	grandParentKey := datastore.IDKey("Society", 5066549580791808, nil)
	grandParentKey.Namespace = "NeverLand"
	parentKey := datastore.IDKey("Family", 5668600916475904, grandParentKey)
	parentKey.Namespace = "NeverLand"

	// parentKey, err := datastore.DecodeKey(*user.ParentID)
	// if err != nil {
	// 	w.Write([]byte("Unable to Decode Key: " + err.Error()))
	// 	appendHeader(w, http.StatusServiceUnavailable)
	// 	return
	// }

	key := datastore.IDKey("Users", 0, parentKey)
	key.Namespace = "NeverLand"

	key, err = client.Put(ctx, key, user)
	if err != nil {
		w.Write([]byte("Failed to create Entity: " + err.Error()))
		appendHeader(w, http.StatusServiceUnavailable)
		return
	}

	user.ParentID = key.String()

	json, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		appendHeader(w, http.StatusNotFound)
		return
	}

	w.Write(json)
	appendHeader(w, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	appendHeader(w, http.StatusOK)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		w.Write([]byte("Failed to create client: " + err.Error()))
		appendHeader(w, http.StatusServiceUnavailable)
		return
	}

	query := datastore.NewQuery("Users").Limit(10).Namespace("NeverLand")

	var users []*User

	keys, err := client.GetAll(ctx, query, &users)
	if err != nil {
		w.Write([]byte("Query Failed: " + err.Error()))
		appendHeader(w, http.StatusNotFound)
		return
	}

	json, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
		appendHeader(w, http.StatusNotFound)
		return
	}
	if len(keys) > 0 {
		w.Write(json)
		appendHeader(w, http.StatusOK)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		w.Write([]byte("Failed to create client: " + err.Error()))
		appendHeader(w, http.StatusServiceUnavailable)
		return
	}

	var user = new(User)

	key, err := datastore.DecodeKey(params["id"])
	if err != nil {
		w.Write([]byte("Unable to Decode Key: " + err.Error()))
		appendHeader(w, http.StatusFailedDependency)
		return
	}

	err = client.Get(ctx, key, user)
	if err != nil {
		w.Write([]byte("Unable to Get: " + err.Error()))
		appendHeader(w, http.StatusNotFound)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("Unable to Marshal: " + err.Error()))
		appendHeader(w, http.StatusServiceUnavailable)
		return
	}
	w.Write(json)
	appendHeader(w, http.StatusOK)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	appendHeader(w, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	appendHeader(w, http.StatusOK)
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {
	appendHeader(w, http.StatusOK)
}
