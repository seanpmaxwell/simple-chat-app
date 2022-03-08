// Store in a single object
const envVars = {
    nodeEnv: process.env.NODE_ENV,
    port: process.env.PORT,
    cookieProps: {
        key: 'simple-chat-app',
        secret: process.env.COOKIE_SECRET,
        name: process.env.COOKIE_NAME,
        options: {
            httpOnly: true,
            signed: true,
            path: (process.env.COOKIE_PATH),
            maxAge: Number(process.env.COOKIE_EXP),
            domain: (process.env.COOKIE_DOMAIN),
            secure: (process.env.SECURE_COOKIE === 'true'),
        },
    },
    jwt: {
        secret: process.env.JWT_SECRET,
        exp: process.env.COOKIE_EXP, // exp at the same time as the cookie
    },
};


// Export default
export default envVars;
