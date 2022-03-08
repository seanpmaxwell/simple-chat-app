/**
 * Setup/start the koa server.
 * 
 * created by Sean Maxwell, 1/28/2022
 */

import Koa from 'koa';
import koaLogger from 'koa-logger';
import json from 'koa-json';
import bodyParser from 'koa-bodyparser';
import session from 'koa-session';

import apiRouter from './routes/api-router';
import staticRouter from './routes/static-router';
import envVars from '@shared/env-vars';


// Constants
const app = new Koa();


// Middlewares
app.use(bodyParser());
app.use(json());
app.use(koaLogger());

// Set keys for cookies (cookies store the jwt)
app.keys = [envVars.cookieProps.name ?? ''];
app.use(session(app));

// Error handling
app.use(async (ctx, next) => {
    try {
        await next();
    } catch (err) {
        ctx.status = (err.status ?? 500);
        ctx.body = {error: err.message};
        ctx.app.emit('error', err, ctx);
    }
});

// API Routes
app.use(apiRouter.routes());
app.use(apiRouter.allowedMethods());

// Static router
app.use(staticRouter.routes());
app.use(staticRouter.allowedMethods());


// Export the server
export default app;
