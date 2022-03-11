import jwt from 'koa-jwt';
import randomstring from 'randomstring';
import envVars from '@shared/env-vars';



/**
 * Parse jsonwebtoken, token data is available at "ctx.state.user".
 */
export function getApiMw() {
    return jwt({
        secret: (envVars.jwt.secret ?? randomstring.generate(100)),
        cookie: envVars.cookieProps.name,
    }).unless({path: [/^\/api\/auth/]});
}


/**
 * Parse jsonwebtoken, token data is available at "ctx.state.user".
 */
export function getSessionMw() {
    return jwt({
        secret: (envVars.jwt.secret ?? randomstring.generate(100)),
        cookie: envVars.cookieProps.name,
        passthrough: true,
    });
}
