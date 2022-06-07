# FeedFish 
![alt text](/assets/feedfish.png)

**Custom pipeline like gitlab ci/cd, github actions**


## Stages

| Name | README |
| ------ | ------ |
| Initial | Prepare the environment stage|
| Processing | Processing stage |
| Proceed | last stage |



## Usage:

### First Step: 
.pipeline.yml  

```yaml
  steps:
    - name: maker
      order: 1
      scripts:
        initial:
          script: "apk add --no-cache git make musl-dev go"
        processing:
          script: "go version"
        proceed: 
          script: "echo 'proceed'"
    
    - name: checker
      order: 2
      scripts:
        initial:
          script: echo 'initial'
        processing:
          script: echo 'processing'
        proceed: 
          script: echo 'proceed'
    
    
```
### Second Step: 
    go run main.go 