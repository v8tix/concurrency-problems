bool mu = false;

proctype proc1() {
    do
    :: atomic {
        if
        :: !mu -> mu = true;
           // Attempting to do work but preempted
           if
           :: !mu -> skip;  // simulate failed check and release
           fi;

           printf("proc1 trying to acquire lock\n");
           mu = false;
        fi;
    }
    od
}

proctype proc2() {
    do
    :: atomic {
        if
        :: !mu -> mu = true;
           // Attempting to do work but preempted
           if
           :: !mu -> skip;  // simulate failed check and release
           fi;
           printf("proc2 trying to acquire lock\n");
           mu = false;
        fi;
    }
    od
}

init {
    atomic {
        run proc1();
        run proc2();
    }
}
