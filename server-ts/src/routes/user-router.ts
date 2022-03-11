import Router from 'koa-router';
import HttpStatusCodes from 'http-status-codes';

import { ParamMissingError } from '@shared/errors';
import userService from '@services/user-service';



// **** Vars/Constants **** //

// Paths
const p = {
    prefix: '/users',
    add: '/',
    fetchAll: '/',
} as const;

// Init router
const router = new Router({prefix: p.prefix});



// **** Functions **** //

/**
 * Add one user.
 */
router.post(p.add, async (ctx) => {
    const { email, name, password } = ctx.request.body;
    // Check params
    if (!email || !name || !password) {
        throw new ParamMissingError();
    }
    await userService.addOne(email, name, password);
    ctx.status = HttpStatusCodes.CREATED;
});


/**
 * Fetch all.
 */
router.get(p.fetchAll, async (ctx) => {
    const users = await userService.fetchAll();
    ctx.body = {users};
});



// Export default
export default router;
