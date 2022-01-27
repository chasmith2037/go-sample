Build docker image

`
docker image build . -t gosample
`

Run image as a container

`
docker run -p 80:80 -p 40003:40003 gosample 
`

List running containers

`
docker ps
`

Stop docker container

`
docker stop <container id>
`

Stop all containers

`
docker ps -a -q | xargs docker stop
`

Remove all containers

`
docker ps -a -q | xargs docker rm
`

Remove all images

`
docker image prune -a
`