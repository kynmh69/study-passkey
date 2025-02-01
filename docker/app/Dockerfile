FROM node:22.13.1

WORKDIR /next/app
COPY ./src/app/package.json .
COPY ./src/app/package-lock.json .

RUN npm i

COPY ./src/app .

EXPOSE 3000

CMD ["npm", "run", "dev"]
