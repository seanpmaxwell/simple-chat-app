import Router from 'koa-router';
import HttpStatusCodes from 'http-status-codes';

import authService from '@services/auth-service';
import { ParamMissingError } from '@shared/errors';
import { tick } from '@shared/functions';
import { getSessionMw } from './middlewares';
import envVars from '@shared/env-vars';


// **** Vars/Constants **** //

// Paths
const p = {
    prefix: '/auth',
    login: '/login',
    sessionData: '/session-data',
    logout: '/logout',
} as const;

// Misc
const router = new Router({prefix: p.prefix});


// **** Functions **** //

/**
 * Login a user by adding a jwt to the cookie.
 */
router.put(p.login, async (ctx) => {
    const { email, password } = ctx.request.body;
    // Check params
    if (!email || !password) {
        throw new ParamMissingError();
    }
    // Check password, add delay if failed
    const { passed, jwt } = await authService.login(email, password);
    ctx.body = {passed};
    if (!passed) {
        await tick(500);
        return;
    }
    // see middlware file
    ctx.cookies.set((envVars.cookieProps.name ?? ''), jwt, envVars.cookieProps.options);
});

/**
 * Logout the user by deleting the cookie.
 */
router.get(p.logout, (ctx) => {
    ctx.cookies.set(envVars.cookieProps.name ?? '');
    ctx.status = HttpStatusCodes.OK;
});

/**
 * Get the logged in user's basic data.
 */
router.get(p.sessionData, getSessionMw(), (ctx) => {
    const data = ctx.state.user;
    ctx.body = {
        id: data ? data.id : -1,
        email: data ? data.email : '',
        name: data ? data.name : '',
        waiting: false,
    };
});


// Export default
export default router;
