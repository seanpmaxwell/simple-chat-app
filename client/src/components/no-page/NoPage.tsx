import { PageWrapper } from "../shared/components";


/**
 * Render this when user inputs a url there is not a page for.
 */
function NoPage(): JSX.Element {
    return (
        <PageWrapper>The requested page was not found</PageWrapper>
    );
}


// Export default
export default NoPage;
