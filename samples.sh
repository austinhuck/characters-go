#!/bin/bash
set -x

host=localhost

curl -X "GET" "http://$host:8000/characters" -H "accept: application/json" --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/characters/1" -H "accept: application/json" --fail-with-body || exit $?
curl -X "PUT" "http://$host:8000/characters/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"id":1,"name":"Grog","description":"Goliath barbarian/fighter ","health":120,"experience":5600}' --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/characters/1" -H "accept: application/json" --fail-with-body || exit $?
curl -X "DELETE" "http://$host:8000/characters/2" -H "accept: application/json" --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/characters/2" -H "accept: application/json"
curl -X "POST" "http://$host:8000/characters" -H "accept: application/json" -H "Content-Type: application/json" -d '{"name":"Grog","description":"Half-elf rogue","health":56,"experience":3400}'
curl -X "POST" "http://$host:8000/characters" -H "accept: application/json" -H "Content-Type: application/json" -d '{"name":"Vax","description":"Half-elf rogue","health":56,"experience":3400}' --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/characters/3" -H "accept: application/json" --fail-with-body || exit $?

curl -X "GET" "http://$host:8000/items" -H "accept: application/json" --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/items/1" -H "accept: application/json" --fail-with-body || exit $?
curl -X "PUT" "http://$host:8000/items/1" -H "accept: application/json" -H "Content-Type: application/json" -d '{"name":"Bag of Holding","description":"A huge bag!","damage":0,"healing":0,"protection":0}' --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/items/1" -H "accept: application/json" --fail-with-body || exit $?
curl -X "POST" "http://$host:8000/items" -H "accept: application/json" -H "Content-Type: application/json" -d '{"name":"Potion of Healing","description":"Heals a number of hitpoints","damage":0,"healing":5,"protection":0}'
curl -X "POST" "http://$host:8000/items" -H "accept: application/json" -H "Content-Type: application/json" -d '{"name":"Potion of Strength","description":"Makes you stronger","damage":0,"healing":0,"protection":2}'
curl -X "GET" "http://$host:8000/items/4" -H "accept: application/json" --fail-with-body || exit $?
curl -X "DELETE" "http://$host:8000/items/1" -H "accept: application/json" --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/items/1" -H "accept: application/json"

curl -X "GET" "http://$host:8000/character/1/inventory" -H "accept: application/json" --fail-with-body || exit $?
curl -X "PATCH" "http://$host:8000/character/1/inventory" -H "accept: application/json" -H "Content-Type: application/json" -d '{"item": 2, "quantity": 1}' --fail-with-body || exit $?
curl -X "PATCH" "http://$host:8000/character/1/inventory" -H "accept: application/json" -H "Content-Type: application/json" -d '{"item": 4, "quantity": 3}' --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/character/1/inventory" -H "accept: application/json" --fail-with-body || exit $?
curl -X "DELETE" "http://$host:8000/character/1/inventory" -H "accept: application/json" --fail-with-body || exit $?
curl -X "GET" "http://$host:8000/character/1/inventory" -H "accept: application/json" --fail-with-body || exit $?
