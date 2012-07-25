package err

const (
	SUCCESS  = 0
	FAILURE  = -1
	CONTINUE = -2 /* iteration has not converged */
	EDOM     = 1  /* input domain error, e.g sqrt(-1) */
	ERANGE   = 2  /* output range error, e.g. exp(1e100) */
	EFAULT   = 3  /* invalid pointer */
	EINVAL   = 4  /* invalid argument supplied by user */
	EFAILED  = 5  /* generic failure */
	EFACTOR  = 6  /* factorization failed */
	ESANITY  = 7  /* sanity check failed - shouldn't happen */
	ENOMEM   = 8  /* malloc failed */
	EBADFUNC = 9  /* problem with user-supplied function */
	ERUNAWAY = 10 /* iterative process is out of control */
	EMAXITER = 11 /* exceeded max number of iterations */
	EZERODIV = 12 /* tried to divide by zero */
	EBADTOL  = 13 /* user specified an invalid tolerance */
	ETOL     = 14 /* failed to reach the specified tolerance */
	EUNDRFLW = 15 /* underflow */
	EOVRFLW  = 16 /* overflow  */
	ELOSS    = 17 /* loss of accuracy */
	EROUND   = 18 /* failed because of roundoff error */
	EBADLEN  = 19 /* matrix, vector lengths are not conformant */
	ENOTSQR  = 20 /* matrix not square */
	ESING    = 21 /* apparent singularity detected */
	EDIVERGE = 22 /* integral or series is divergent */
	EUNSUP   = 23 /* requested feature is not supported by the hardware */
	EUNIMPL  = 24 /* requested feature not (yet) implemented */
	ECACHE   = 25 /* cache limit exceeded */
	ETABLE   = 26 /* table limit exceeded */
	ENOPROG  = 27 /* iteration is not making progress towards solution */
	ENOPROGJ = 28 /* jacobian evaluations are not improving the solution */
	ETOLF    = 29 /* cannot reach the specified tolerance in F */
	ETOLX    = 30 /* cannot reach the specified tolerance in X */
	ETOLG    = 31 /* cannot reach the specified tolerance in gradient */
	EOF      = 32 /* end of file */
)
