import react from 'react';

import {
    Grid,
    Paper,
    styled,
} from '@mui/material';

import { PageWrapper } from '../shared/components';
import LoginForm from './LoginForm';


interface IProps {
    fetchSessionData: () => Promise<void>;
}

/**
 * Home.
 */
function Home(props: IProps): JSX.Element {
    return (<PageWrapper>
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
                <LoginPaper elevation={12}>
                    <LoginForm fetchSessionData={() => props.fetchSessionData()}/>
                </LoginPaper>
            </Grid>
        </Grid>
    </PageWrapper>);
}

const LoginPaper = styled(Paper)(({ theme }) => ({
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
    backgroundColor: '#fafafa',
}));


// Export default.
export default Home;
