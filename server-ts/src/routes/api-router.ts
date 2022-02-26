import Router from 'koa-router';
import authRouter from './auth-router';
import userRouter from './user-router';
import { getApiMw } from './middlewares';


// Init router
const router = new Router({prefix: '/api'});

// Routes below this should be protected
router.use(getApiMw());


// auth router
router.use(authRouter.routes());
router.use(authRouter.allowedMethods());

// user router
router.use(userRouter.routes());
router.use(userRouter.allowedMethods());


// Export default
export default router;
