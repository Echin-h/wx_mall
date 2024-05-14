FROM golang:latest AS builder

COPY . /build

WORKDIR /build

