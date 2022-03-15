import { useDispatch, useSelector } from 'react-redux';
import { Grid } from '@mui/material';

import { PageWrapper } from '../../shared/components/PageWrapper';
import { useSetState } from '../../shared/hooks';
import { actions as slc, IReduxSliceState } from './ReduxSlice';
import ReduxSlice from './ReduxChild/ReduxChild';


// **** Types **** //

interface IState {
    count: number;
}


// **** Redux **** //

/**
 * Render()
 */
function Redux(): JSX.Element {
    const [state, setState ] = useSetState(init());
    const dispatch = useDispatch();
    const sliceState = useSelector<IReduxSliceState, IReduxSliceState>(sliceState => sliceState);
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
                <div>
                    Parent Local Count: {state.count}
                </div>
                <div>
                    <button onClick={() => setState({count: state.count + 1})}>
                        Incremement
                    </button>
                </div>
                <div>
                    <button onClick={() => setState({count: state.count - 1})}>
                        Decrement
                    </button>
                </div>
                <div>
                    <button onClick={() => setState({count: 0})}>
                        Reset
                    </button>
                </div>
                &nbsp;
                <div>
                    Parent Slice Count: {sliceState.count}
                </div>
                <div>
                    <button onClick={() => dispatch(slc.increment())}>
                        Incremement
                    </button>
                </div>
                <div>
                    <button onClick={() => dispatch(slc.decrement())}>
                        Decrement
                    </button>
                </div>
                <div>
                    <button onClick={() => dispatch(slc.reset())}>
                        Reset
                    </button>
                </div>
                &nbsp;
                <ReduxSlice/>
            </Grid>
        </Grid>
    </PageWrapper>);
}


/**** Functions ****/

/**
 * Get the initial state.
 */
function init(): IState {
    return {
        count: 0,
    };
}


// Export default.
export default Redux;
