# About The Project
Example golang using gin framework everything you need :)

## Installation:
   ### Prerequisites
   - docker
      for download docker in [link](https://docs.docker.com/engine/install/)

   - docker-compose
      for download docker in [link](https://docs.docker.com/compose/install/)
  
   ### create volume

    docker volume create postgres_data
   
   ### create network
   
    docker network create main
   
   ### run service 
 
    docker-compose up -d --build
