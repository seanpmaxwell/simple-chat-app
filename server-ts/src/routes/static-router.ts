import Router from 'koa-router';


// **** Vals **** //

// Misc
const router = new Router();


// **** Functions **** //

/**
 * Serve html.
 */
router.get('', (ctx) => {
    ctx.type = 'html';
    ctx.body = '<div>Hello, you are running Koa in development mode.</div>';
});


// Export default
export default router;
