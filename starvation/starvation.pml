bool mu = false;

proctype proc1() {
    do
    :: atomic {
        !mu -> mu = true;
        printf("proc1 executing\n");
        mu = false;
    }
    od
}

proctype proc2() {
    do
    :: atomic {
        !mu -> mu = true;
        printf("proc2 executing\n");
        mu = false;
    }
    od
}

proctype proc3() {
    do
    :: atomic {
        printf("proc3 waiting to execute\n");
        !mu -> mu = true;
        printf("proc3 executing and holding the lock\n");
        // Simulate holding the lock for an extended period
        int i;
        i = 0;
        do
        :: (i < 1000) -> i++;
        :: else -> break;
        od;
        printf("proc3 releasing the lock\n");
        mu = false;
    }
    od
}

init {
    run proc1();
    run proc2();
    run proc3();
}
