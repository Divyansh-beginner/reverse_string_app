//check_redis.go
package main
import (
"fmt"
"log"
"net/http"
"time"
)

func check_function(w http.ResponseWriter , r *http.Request) {

	formerror := r.ParseForm()
	if formerror != nil{
	http.Error(w , "failed to parse the data", http.StatusBadRequest)
	return
	}

	input := r.FormValue("text")
	if input == ""{
	http.Error(w , "no text field ",http.StatusBadRequest)
	return
	}
	setupredisclient()
	defer redisclient.Close()
	redispingerr := checkredisconnection()
	if redispingerr != nil{
		log.Println(redispingerr)
		http.Error(w , redispingerr.Error() , http.StatusInternalServerError)
		connectionerr := connectandcheckdbconnection()
		if connectionerr != nil {
		log.Println(connectionerr)
		http.Error(w , connectionerr.Error() , http.StatusInternalServerError)
		return
		}
		defer db.Close()
		exists ,dbvalue , dberr := checkindb(input)
		if exists {
		fmt.Fprintln(w,dbvalue)
		return
		}
		if dberr != nil {
		log.Println(dberr)
		http.Error(w , dberr.Error() ,  http.StatusInternalServerError)
		return
		}
		fmt.Fprintln(w , "not found")
		return
	}else{
		cachevalue, cacheerror := getfromcache(input)
		if cacheerror != nil {
			log.Println(cacheerror)
			http.Error(w , cacheerror.Error() , http.StatusInternalServerError)
		}else if cachevalue != nil {
			fmt.Fprintln(w,cachevalue)
			return
		}
		if cachevalue ==nil{
			connectionerr := connectandcheckdbconnection()
			if connectionerr != nil {
				log.Println(connectionerr)
				http.Error(w , connectionerr.Error() , http.StatusInternalServerError)
				return
			}
			defer db.Close()
			exists ,dbvalue , dberr := checkindb(input)
			if dberr != nil {
				log.Println("database lookup failed ! error: %v",dberr)
				http.Error(w ,dberr.Error(),http.StatusInternalServerError)
				return
			}
			if exists {
				fmt.Fprintln(w,dbvalue)
				setcacheerror := settocache(input , dbvalue , 10*time.Minute)
				if setcacheerror != nil{
					log.Println(setcacheerror)
					http.Error(w,setcacheerror.Error(),http.StatusInternalServerError)
				}
				return
			}
			fmt.Fprintln(w,"not found ")
			return
		}
	}
}