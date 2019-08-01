## Simple alert reminder

Do you need to do something in an hour and a half minutes and wan't to be reminded while you use your pc?

```
go install github.com\j4rv\gostuff\remindme  
remindme -h 1 -m 30 -msg "Remember to walk the dog!"
```

Another example:

```
remindme -m 6 -msg "Go check the cofeemaker!"
```

The alert will be a shown using the *github.com/gen2brain/beeep* module.  
**(Closing the console will stop the app and the alert will not show!)**