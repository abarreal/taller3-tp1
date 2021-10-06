// The following script uses local caching to give the illusion
// of a strongly consistent visit counter when there is in fact eventual 
// consistency. It assumes that the counter retrieved from the server 
// is at least one hit below the eventual value.
const SERVER_COUNTER = 'SERVER_COUNTER';
const CLIENT_COUNTER = 'CLIENT_COUNTER';

// Get a reference to the visit counter provided by the server.
let visitCounter = document.getElementById('visitcounter');
let serverCounter = visitCounter.textContent;

if (serverCounter && serverCounter.length > 0 && serverCounter != "-1") {
    serverCounter = parseInt(serverCounter);

    // Get the amount of visits reported by the server, the last time
    // we accessed the page.
    let cachedServerCounter = sessionStorage.getItem(SERVER_COUNTER);
    let clientCounter = sessionStorage.getItem(CLIENT_COUNTER);

    // Define a flag to tell whether we should reset the local counters.
    let reset = false;

    if (!cachedServerCounter) {
        reset = true;
    } else {
        // We do have data in the storage, so we parse it to integers.
        cachedServerCounter = parseInt(cachedServerCounter);
        clientCounter = parseInt(clientCounter);
        // If the server reported a visit count that is greater from
        // what we have in cache, then we use the server's counter
        // rather than our local version.
        reset = (serverCounter > (cachedServerCounter+clientCounter));
    }
   
    if (reset) {
        // The server reported a new value for its counter, so we update
        // the locally cached version of that counter and also reset
        // the local client counter to 0. We assume that the server,
        // even if handling this counter in an eventually consistent way,
        // did take into account this visit when reporting the count.
        sessionStorage.setItem(SERVER_COUNTER, serverCounter);
        sessionStorage.setItem(CLIENT_COUNTER, 0);
    } else {
        // Increase the local counter by one for the browser to keep
        // it's own count until the server reports an updated counter.
        clientCounter += 1;
        sessionStorage.setItem(CLIENT_COUNTER, clientCounter);
    }

    // Parse the client counter to ensure that we have an integer.
    clientCounter = parseInt(sessionStorage.getItem(CLIENT_COUNTER));
    // Update the HTML counter and make visible.
    visitCounter.textContent = serverCounter + clientCounter;
    visitCounter.removeAttribute('hidden');
}