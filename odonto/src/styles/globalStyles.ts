import { createGlobalStyle } from 'styled-components';

export default createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    outline: 0;
    transition: 0.1s;

    font-family: 'Open Sans', sans-serif;
  }

  body {
    background-color: var(--white);
  }

  :root {
    --primary: #3949AB;
    --danger: red;
    --primary-light-color: #6F74DD;
    --primary-text: #343434;
    --link-text: #6F74DD;
    --primary-text-light: #828282;
    --control-border: #D0D0D0;
    --light-gray: #F5F5F5;
    
    --white: #FFFFFF;
    --gray: #D0D0D0;


    --radius-default: 0.5rem; // 8px
  }
`;
