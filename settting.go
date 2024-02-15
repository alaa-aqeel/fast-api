package main

import "time"

// APP
const APP_URL = "0.0.0.0:8080"

// JWT
var SECRET_KEY []byte = []byte("107684d9d2b8664121807f3d654")
var EXPIRE_TOKEN = time.Now().Add(time.Hour * 24).Unix()
