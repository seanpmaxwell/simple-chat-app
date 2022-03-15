import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import Grid from '@mui/material/Grid';
import { TextField } from '@mui/material';

import { useSetState } from '../../shared/hooks';
import usersHttp from '../../shared/http/users-http';


// **** Vars ****//

export enum Modes {
    Closed = 1,
    Add,
    Edit,
}


// **** Types ****//

// UserForm Props
interface IProps {
    mode: Modes;
    setMode: (mode: Modes) => void;
}

// UserForm State
interface IState {
    email: string;
    name: string;
    password: string
    addUserFailed: boolean
}


// **** UserForm **** //

/**
 * Render()
 */
function UserForm(props: IProps): JSX.Element {
    const [state, setState ] = useSetState(init());
    return (
        <Dialog
            aria-labelledby='user-form'
            aria-describedby='user-form-description'
            maxWidth='xs'
            fullWidth={true}
            open={props.mode !== Modes.Closed}
            onClose={() =>  props.setMode(Modes.Closed)}
        >
            <DialogTitle id='user-form'>
                <span className='flex-start-row'>
                    {props.mode === Modes.Add ? 'Add' : 'Edit'} User
                </span>
            </DialogTitle>

            <DialogContent>
                <Grid
                    container={true}
                    direction="column"
                    justifyContent="center"
                    alignItems="stretch"
                    p={2}
                >
                    <Grid item={true}>
                        <TextField
                            id="user-add-email"
                            label="Email"
                            type="email"
                            variant="outlined"
                            fullWidth={true}
                            value={state.email}
                            onChange={e => setState({email: e.currentTarget.value})}
                        />
                    </Grid>
                    <br/>
                    <Grid>
                        <TextField
                            id="user-add-ename"
                            label="Name"
                            type="text"
                            variant="outlined"
                            fullWidth={true}
                            value={state.name}
                            onChange={e => setState({name: e.currentTarget.value})}
                        />
                    </Grid>
                    <br/>
                    <Grid>
                        <TextField
                            id="user-add-name"
                            label="Password"
                            type="password"
                            variant="outlined"
                            fullWidth={true}
                            value={state.password}
                            onChange={e => setState({password: e.currentTarget.value})}
                        />
                    </Grid>
                </Grid>
            </DialogContent>

            <DialogActions>
                <Button
                    color='primary'
                    variant='contained'
                    disabled={!state.email || !state.name || !state.password}
                    onClick={() => addUser(state).then(passed => {
                        if (passed) {
                            setState(init());
                            props.setMode(Modes.Closed);
                        } else {
                            setState({addUserFailed: true});
                        }
                    })}
                >
                    Submit
                </Button>
                <Button
                    color='error'
                    variant='contained'
                    onClick={() => {
                        setState(init());
                        props.setMode(Modes.Closed);
                    }}
                >
                    Cancel
                </Button>
            </DialogActions>
        </Dialog>
    );
}


// **** Functions **** //

/**
 * Get Initial State
 */
function init(): IState {
    return {
        email: '',
        name: '',
        password: '',
        addUserFailed: false,
    };
}

/**
 * Add a new user
 */
async function addUser(state: IState): Promise<boolean> {
    const { email, name, password } = state;
    let success = false;
    try {
        await usersHttp.add(email, name, password);
        success = true;
    } catch (err) {
        console.error(err);
    }
    return success;
} 


// Export default
export default UserForm;
