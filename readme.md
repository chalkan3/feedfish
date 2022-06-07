# FeedFish 
![alt text](/assets/feedfish.png)

**Custom pipeline like gitlab ci/cd, github actions**

## Usage:

### First Step: 
Fill the .pipeline.yml

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