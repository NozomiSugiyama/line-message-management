declare const process: {
    env: {
        [key: string]: string | undefined
    }
};

export default {
    api: {
        uri: process.env.API_URI
    }
};
