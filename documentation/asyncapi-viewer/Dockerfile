# syntax=docker/dockerfile:1

FROM node:18

WORKDIR /app

COPY package.json package-lock.json ./

RUN --mount=type=cache,target=/root/.npm npm install --force

COPY . .