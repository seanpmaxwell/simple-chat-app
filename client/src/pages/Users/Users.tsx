import { Button, Grid } from '@mui/material';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

import { PageWrapper } from '../../shared/components/PageWrapper';
import { useSetState } from '../../shared/hooks';
import UserForm, { Modes } from './UserForm';
import usersHttp, { IUser } from '../../shared/http/users-http';
import { useEffect } from 'react';


// **** Types **** //

// Users State
interface IState {
    userFormMode: Modes;
    users: IUser[];
}


// **** Users **** //

/**
 * Render()
 */
function Users(): JSX.Element {
    const [state, setState ] = useSetState(init());
    useEffect(() => {
        fetchUsers().then(users => setState({users}));
    }, [setState]);
    // Return
    return (<PageWrapper>
        <br/>
        <Button
            color='primary'
            variant='contained'
            onClick={() => setState({userFormMode: Modes.Add})}
        >
            Users
        </Button>
        <Grid
            container={true}
            direction="column"
            justifyContent="center"
            alignItems="center"
            sx={{
                flexGrow: 1,
            }}
        >
            <Grid
                item={true}
                lg={true}
                sx={{
                    marginTop: '15%',
                }}
            >
                 <TableContainer component={Paper}>
                    <Table
                        aria-label="simple table"
                        sx={{ minWidth: 650 }}
                    >
                        <TableHead>
                            <TableRow>
                                <TableCell sx={{fontWeight: 'bold'}}>Email</TableCell>
                                <TableCell sx={{fontWeight: 'bold'}}>Name</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {state.users.map((user) => (
                                <TableRow key={user.id}>
                                    <TableCell>{user.email}</TableCell>
                                    <TableCell>{user.name}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                 </TableContainer>
            </Grid>
        </Grid>
        <UserForm
            mode={state.userFormMode}
            setMode={userFormMode => setState({userFormMode})}
        />
    </PageWrapper>);
}


// **** Functions **** //

/**
 * Get initial State
 */
function init(): IState {
    return {
        userFormMode: Modes.Closed,
        users: [],
    };
}

/**
 * Fetch all users from the database
 */
async function fetchUsers(): Promise<IUser[]> {
    let users: IUser[] = [];
    try {
        const resp = await usersHttp.fetchAll();
        users = resp.users;
    } catch (err) {
        console.error(err);
    }
    return users;
}


// Export default
export default Users;
