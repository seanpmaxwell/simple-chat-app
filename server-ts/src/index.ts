import '../env';
import logger from 'jet-logger';
import app from './server';


// Constants
const serverStartMsg = 'Koa Started on port localhost:',
    port = process.env.PORT;

    
// Start server
app.listen(port, () => {
    logger.imp(serverStartMsg + port);
});
