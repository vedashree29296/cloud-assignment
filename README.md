# cloud-assignment
Cloud Assignment to create API to find and append copyright symbol to company names

## The entire Documentation can be found at:
https://github.com/vedashree29296/cloud-assignment/wiki

Links to Pages:

### [1. Overview](https://github.com/vedashree29296/cloud-assignment/wiki/1.-Overview)

### [2. Try out the API](https://github.com/vedashree29296/cloud-assignment/wiki/2.-Try-out-the-API)

### [3. API & Development Specifics](https://github.com/vedashree29296/cloud-assignment/wiki/3.-API-&-Development-Specifics)

### [4. Infrastructure and Deployments](https://github.com/vedashree29296/cloud-assignment/wiki/4.-Infrastructure-and-Deployment)

## TL;DR

#### Swagger UI and API Documentation is available at:

[http://swagger-ui-cloud-assignment.s3-website-us-east-1.amazonaws.com/](http://swagger-ui-cloud-assignment.s3-website-us-east-1.amazonaws.com/)

#### API endpoint:

[https://kepxuc1vi5.execute-api.us-east-1.amazonaws.com/dev/replaceOrganisation](https://kepxuc1vi5.execute-api.us-east-1.amazonaws.com/dev/replaceOrganisation)

#### Documentation 
[https://github.com/vedashree29296/cloud-assignment/wiki](https://github.com/vedashree29296/cloud-assignment/wiki)

#### Sample curl command:

```curl -X POST "https://kepxuc1vi5.execute-api.us-east-1.amazonaws.com/dev/replaceOrganisation" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{\"text\":\"Facebook, Netflix, Amazon and Google are called the FANG companies.\",\"add_organisation\":[\"Facebook\"]}"```

### Assignment Specifics:
- Developed in Go
- Deployed on AWS Lambda + API Gateway using Serverless Framework.
- CI/CD using GitHub actions