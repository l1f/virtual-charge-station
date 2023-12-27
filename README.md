# virtual-charge-station

## About the Projekt

In the last year and a half, I have been working a lot on the server-side implementation of OCPP. 
In doing so, I have repeatedly encountered the problem that I lack the possibility to start one or more virtual charging stations on demand 
in order to be able to carry out tests quickly. 

This meant that I needed a physical charging station at all times to test whether my server was working properly. 
- Which made it difficult to work from anywhere that wasn't my desk.

So here's the next project, which I'm going to work on very intensively for two weeks and then probably won't touch for 
a very long time because my focus will switch to something else.

## Feature roadmap

- [ ] Interface for creating, deleting and updating stations
- [ ] Basic OCPP 1.6 support (Sending and receiving data)
- [ ] OCPP 1.6 feature sets
  - [ ] Core
  - [ ] Firmware Managment
  - [ ] Local Auth
  - [ ] Reservation
  - [ ] Smart Charging
  - [ ] Remote Trigger
- [ ] CI Mode

### Hardware Support

A very cool but in my eyes unrelated goal of mine is to also support hardware like RFID, Girocard and credit card readers to test payment functions. 
But I'm not even including that in the roadmap, because I don't think I have that much time to work on a toy project. 

### Signed meter values

Another even more unrealistic goal of mine is that the OCPP stations that are created with the tool also have the option to generate and transmit signed meter values. 

## The stack

And because this is a toy project, I'm testing technologies that I've wanted to play around with for a long time. 

- [Go](https://go.dev/) (Not necessarily a technology that I don't know yet)
- [templ](https://templ.guide)
- [htmx](https://htmx.org/)
- [chi](https://go-chi.io/#/)
- [urfave/cli](https://cli.urfave.org/)
- [zap](https://github.com/uber-go/zap)
- [air](https://github.com/cosmtrek/air)

## Development

To work on the project, you currently need a current Go compiler, templ and git. 

Optionally, you can install and use air. 
