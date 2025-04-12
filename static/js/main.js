document.addEventListener('DOMContentLoaded', function() {
    // Add event listeners for form submission
    const form = document.querySelector('.input-form');
    if (form) {
        form.addEventListener('submit', function(e) {
            const input = form.querySelector('input[name="input"]').value;
            const frames = form.querySelector('input[name="frames"]').value;
            const algorithm = form.querySelector('select[name="algorithm"]').value;

            // Validate input
            if (!input || !frames || !algorithm) {
                e.preventDefault();
                alert('Please fill in all fields');
                return;
            }

            // Validate page sequence format
            const pages = input.trim().split(/\s+/);
            if (!pages.every(page => !isNaN(page))) {
                e.preventDefault();
                alert('Please enter valid page numbers separated by spaces');
                return;
            }
        });
    }

    // Add event listeners for buttons
    const buttons = document.querySelectorAll('button');
    buttons.forEach(button => {
        button.addEventListener('click', function() {
            if (this.disabled) {
                return;
            }
            
            // Add visual feedback
            this.style.opacity = '0.7';
            setTimeout(() => {
                this.style.opacity = '1';
            }, 200);
        });
    });
}); 