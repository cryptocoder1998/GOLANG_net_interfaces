package main

import "testing"

func TestIncorrectIP_Version(t *testing.T) {
	str, err := SendVersionRequest("test_ip", "8000")
	if str != "Server is unavailable" {
		t.Error("Expected str: Server is unavailable")
	}
	if err == nil {
		t.Error("Expected an err returned")
	}
}
func TestIncorrectPort_Version(t *testing.T) {
	str, err := SendVersionRequest("127.0.0.1", "test_port")
	if str != "Server is unavailable" {
		t.Error("Expected str: Server is unavailable")
	}
	if err == nil {
		t.Error("Expected an err returned")
	}
}
func TestIncorrectIP_Enum(t *testing.T) {
	str, err, _ := SendEnumerateRequest("test_ip", "8000")
	if str != "Server is unavailable" {
		t.Error("Expected str: Server is unavailable")
	}
	if err == nil {
		t.Error("Expected an err returned")
	}
}
func TestIncorrectPort_Enum(t *testing.T) {
	str, err, _ := SendEnumerateRequest("127.0.0.1", "test_port")
	if str != "Server is unavailable" {
		t.Error("Expected str: Server is unavailable")
	}
	if err == nil {
		t.Error("Expected an err returned")
	}
}
func TestIncorrectIP_IntInfo(t *testing.T) {
	str, err, _ := SendIntRequest("test_ip", "8000", "test_name")
	if str != "Server is unavailable" {
		t.Error("Expected str: Server is unavailable")
	}
	if err == nil {
		t.Error("Expected an err returned")
	}
}
func TestIncorrectPort_IntInfo(t *testing.T) {
	str, err, _ := SendIntRequest("127.0.0.1", "test_port", "test_name")
	if str != "Server is unavailable" {
		t.Error("Expected str: Server is unavailable")
	}
	if err == nil {
		t.Error("Expected an err returned")
	}
}
func TestIncorrectServerInterface(t *testing.T) {
	str, _, _ := SendIntRequest("127.0.0.1", "8000", "test_int")
	if str != "Such interface doesn't exist" {
		t.Error("Expected str: Such interface doens't exist")
	}
}
