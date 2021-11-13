db.createUser({
    user: 'pass',
    pwd: 'user',
    roles: [
        {
            role: 'readWrite',
            db: 'spectra',
        },
    ],
});

db = new Mongo().getDB("spectra");

db.createCollection('spectra_request', { capped: false });