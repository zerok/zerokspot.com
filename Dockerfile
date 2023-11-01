# Images used for Dagger
FROM zerok/webmentiond:latest
FROM alpine:3.18
FROM ubuntu:22.04
FROM golang:1.21.3 as builder
FROM klakegg/hugo:0.111.3-ext-ubuntu
