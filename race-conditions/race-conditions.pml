int balance = 250;

proctype incrementBalance() {
    do
    :: balance = balance + 10;
    od
}

proctype decrementBalance() {
    do
    :: balance = balance - 10;
    od
}

proctype monitorBalance() {
    do
    ::
	printf("Balance changed: %d\n", balance);
	assert(balance == 250); // This will cause the model checker to stop
    od
}

init {
    run incrementBalance();
    run decrementBalance();
    run monitorBalance();
}