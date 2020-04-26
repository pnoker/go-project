/*
 * Copyright Pnoker. All Rights Reserved.
 */

package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

const (
	str string = "0123456789abcdefghijklmnopqrstuvwxyz"
)

// timer
func Timer(execute func(), second int) *time.Ticker {
	ticker := time.NewTicker(time.Second * time.Duration(second))
	go func(ticker *time.Ticker) {
		for range ticker.C {
			execute()
		}
	}(ticker)
	return ticker
}

// float32 convert to string
func FloatToString(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 6, 64)
}

// get string md5 hash
func HashString(data string) string {
	hash := md5.New()
	io.WriteString(hash, data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// random string
func RandomString(length int) string {
	var result []byte
	bytes := []byte(str)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[random.Intn(len(bytes))])
	}
	return string(result)
}

// random float
func RandomFloat(min, max float32) float32 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float32()*(max-min) + min
	return random
}

// random walk float
func RandomWalkFloat(base, min, max, stepMin, stepMax float32) float32 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float32()*(stepMax-stepMin) + stepMin + base
	if random < min || random > max {
		return RandomWalkFloat(base, min, max, stepMin, stepMax)
	}
	return random
}
