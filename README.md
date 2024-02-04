# uhash

Perform hashes of files, streams, etc in parallel.

    echo 'hello world' | sha256sum -

Becomes

    echo 'hello world' | uhash -h sha256 -
