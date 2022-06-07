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
          script: "docker ps && sleep 10"
        processing:
          script: "go build -o frompipeline cmd/main.go"
        proceed: 
          script: ls
    
    - name: updater
      order: 2
      scripts:
        initial:
          script: ls
        processing:
          script: ls -la
        proceed: 
          script: ls

    - name: char
      order: 3
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