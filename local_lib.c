/*
functions in this file are related to interruption and system signals.

More here : https://github.com/polarspetroll/gopio/tree/main/examples/interruption
*/

#include <stdlib.h>
#include <signal.h>
#include <stdbool.h>
#include <wiringPi.h>


void close_pin(int pnum, int value) {
  bool v = true;

  void close(int i) {
    if (value == 0) {
      digitalWrite(pnum, LOW);
    }else {
      digitalWrite(pnum, HIGH);
    }
    v = false;
    exit(0);
  }

  signal(SIGHUP, close);
	signal(SIGTERM, close); // listens for system signals
  signal(SIGINT, close);

  while (v) {
  }

}
