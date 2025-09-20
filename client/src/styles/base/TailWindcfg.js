 tailwind.config = {
            theme: {
                extend: {
                    keyfreames: {
                        gradient: {
                            '0%, 100%': { 'background-position': '0% 50%' },
                            '50%': { 'background-position': '100% 50%' },
                            '100%': { 'background-position': '0% 50%' },
                        },
                    },
                    animation: {
                        gradient: 'gradient 3s ease infinite',
                    },
                    colors: {
                        'primary': '#ff3333',
                        'primary-dark': '#cc0000',
                        'primary-darker': '#221f1fff',

                        'secondary-dark': '#333030ff',
                        'secondary-darker': '#2b2727ff',

                        'dark': '#131212ff',
                        'darker': '#0a0a0a',
                        'light': '#f8f8f8',
                    },
                    fontFamily: {
                        'code': ['Consolas', 'Monaco', 'Courier New', 'monospace'],
                        'sans': ['Inter', 'system-ui', 'sans-serif'],
                    }
                }
            }
        }