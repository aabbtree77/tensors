<div style="text-align:center; font-size: 2.5rem; margin-top: 3rem; margin-bottom: 1.6rem;"> 
    <div style="font-size: 2rem; margin-top: 2rem; margin-bottom: 2rem;">Notes on <a href="https://aapt.scitation.org/doi/10.1119/1.1976018">Shankland (1970)</a></div>
    <a style="font-size: 1.5rem; margin-top: 1.5rem; margin-bottom: 1.5rem;" href="https://aabbtree77.github.io/">Ramūnas Girdziušas</a>
    <div style="font-size: 1rem; margin-top: 1rem; margin-bottom: 1rem;">Last Update: September 26, 2024</div>
</div>

## Tensor Freedom

A tensor is anything that transforms like a tensor.

**Exercise 1.** Show that tensors allow the following index positioning freedom:

$$
X{^i}{^j}{_k} = X{^i}{_k}{^j} = X{_k}{^i}{^j}.
$$

Hint: Define a tensor as a weighted sum of the Kronecker products of the basis vectors:

$$
\begin{align}
X{^i}{^j}{_k}\equiv \sum_{ijk}\,c_{ijk}\,e^{i}\otimes e^{j}\otimes e_{k}\,,
\end{align}
$$

where $c_{ijk}$ are some constants. Assume that $e^{i}$ is a unit row-vector, and $e_{k}$ is a unit column-vector. Note the property:

$$
\begin{align}
a^{i}\otimes b^{j} &\neq b^{j}\otimes a^{i}\,,\\
a^{i}\otimes b_{j} &= b_{j}\otimes a^{i}\,.
\end{align}
$$

It holds for any row-vector $a^{i}$ and column-vector $b_{j}$.

A good book on tensors is [Schaum's Tensor Calculus](https://kishorekoduvayur.wordpress.com/wp-content/uploads/2017/12/schaums-tensor-calculus-238.pdf) by David C. Kay. Sadly, we do not have anything like it written for spinors and Lie groups.

## Shankland's Tensor Algebras

A problem that ChatGPT cannot solve: Given the [four-vector](https://en.wikipedia.org/wiki/Four-vector) $k_{\mu}$ and [the metric tensor](https://en.wikipedia.org/wiki/Metric_tensor) $g_{\mu\nu}$, write down the most general dimensionless tensor ${T_{\mu\nu}}^{\rho \sigma}$ symmetric under the permutations of the covariant indices $(\mu, \nu)$, and also symmetric w.r.t. the permutations of contravariant indices $(\rho, \sigma)$. It should be a sum of linearly independent terms, each with a manifest symmetry, and at most fourth order in $k_{\mu}$.

[Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) jumps into the answer, which is a linear combination of

$$
\begin{align}
X_{1} & = \frac{1}{4} {g_{(\mu}}^{(\rho} {g_{\nu)}}^{\sigma)}\,,\\
X_{2} & = g_{\mu \nu} g^{\rho \sigma}\,, \\
X_{3} & = \frac{1}{k^2} g_{\mu \nu} k^{\rho} k^{\sigma}\,, \\
X_{4} & = \frac{1}{k^2} k_{\mu} k_{\nu} g^{\rho \sigma}\,, \\
X_{5} & = \frac{1}{k^4} k_{\mu} k_{\nu} k^{\rho} k^{\sigma}\,, \\
X_{6} & = \frac{1}{k^2} k_{(\mu} {g_{\nu)}}^{(\rho} k^{\sigma)}\,.
\end{align}
$$

Here the division by 4 of the first basis element is not important, but it turns $X_{1}$ into an index symmetrization operator if one acts with it on any two-index tensor.

**Magically, these tensor basis expressions form an algebra under the product defined as**

$$
A\, B \equiv {A_{\mu \nu}}^{\alpha \beta} {B_{\alpha \beta}}^{\rho \sigma}\,,
$$

with the implied summation over repeated indices.

**Exercise 2.** Show that

$$
X_{6} X_{6} = 8 X_{5} + 2 X_{6},
$$

which is not $8 X_{4} + 2 X_{6}$ stated by [Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018).

Hint:

The definitions of the product and index symmetrization $(\cdot, \cdot)$ lead to $X_{6} X_{6}$ expressed as

$$
\begin{align}
& \equiv \frac{1}{k^2} k_{(\mu} {g_{\nu)}}^{(\alpha} k^{\beta)} \frac{1}{k^2} k_{(\alpha} {g_{\beta)}}^{(\rho} k^{\sigma)}\\
 &= \frac{1}{k^4} \Big( k_{\mu} {g_{\nu}}^{\alpha} k^{\beta} + k_{\mu} {g_{\nu}}^{\beta} k^{\alpha} \\
            & \;\;\;\;+ k_{\nu} {g_{\mu}}^{\alpha} k^{\beta} + k_{\nu} {g_{\mu}}^{\beta} k^{\alpha} \Big)\\
            & \;\;\;\;\cdot \Big( k_{\alpha} {g_{\beta}}^{\rho} k^{\sigma} + k_{\alpha} {g_{\beta}}^{\sigma} k^{\rho} \\
            & \;\;\;\;+ k_{\beta} {g_{\alpha}}^{\rho} k^{\sigma} + k_{\beta} {g_{\alpha}}^{\sigma} k^{\rho} \Big).
\end{align}
$$

This will produce 16 terms, but they will further simplify with the use of

$$
\begin{align}
k_{\alpha} k^{\alpha} & = k^2 \\
{g_{\nu}}^{\alpha}k_{\alpha} &= k_{\nu} \\
{g_{\nu}}^{\beta} {g_{\beta}}^{\sigma} &= {g_{\nu}}^{\sigma}\;,
\end{align}
$$

e.g.

$$
\begin{align}
t_{1} & = k_{\mu} {g_{\nu}}^{\alpha} k^{\beta} \cdot k_{\alpha} {g_{\beta}}^{\rho} k^{\sigma} \\
& = k_{\mu} \big({g_{\nu}}^{\alpha} k_{\alpha}\big) \big({g_{\beta}}^{\rho} k^{\beta}\big) k^{\sigma} \\
&= k_{\mu} k_{\nu} k^{\rho} k^{\sigma}.
\end{align}
$$

A page later ;) there will be 8 such terms and a doubled $X_{6}$ left. 

<div class="imgcontainer">
<img src="imgs/lifetime.jpg" title="Volumetric Lighting in Max Payne (2008)" width=100%>
</div>

## The Spectrum of a Tensor Field

The products $X_{i}X_{j}$ and the traces $\textit{tr} X_{i}$ determine the spectrum of the algebra. Shankland does not define a tensor field, but it is assumed that $X_{i}$ will form an operator applied to build a quadratic form for a field, where $k_{i}$ becomes a four-nabla. The PhD thesis of [K.J. Barnes (1963)](https://spiral.imperial.ac.uk/bitstream/10044/1/13413/2/Barnes-KJ-1963-PhD-Thesis.pdf) sheds more light here, but his notation requires getting used to.

The traces are defined as

$$
\begin{align}
\textit{tr}\,{X_{\mu \nu}}^{\rho \sigma}  \equiv {X_{\mu \nu}}^{\mu \nu}\,.
\end{align}
$$

They demand the values of the metric tensor $g_{ij}$, along with the mixed tensor

$$
\begin{align}
{g_{i}}^{j}={\delta_{i}}^{j}=\begin{cases}
0 & i \ne j \\
1 & i = j
\end{cases}\;\;.
\end{align}
$$

There is no need to know these values when getting the product tables $X_{i}X_{j}$. 

[Shankland (1970)](https://aapt.scitation.org/doi/10.1119/1.1976018) applies [the Faddeev - LeVerrier algorithm](https://en.wikipedia.org/wiki/Faddeev%E2%80%93LeVerrier_algorithm), or rather its advanced variant extended to tackle multiple eigenvalues, see e.g. [Helmberg and Wagner (1993)](https://core.ac.uk/download/pdf/81192811.pdf). Here there are no eigenvectors in a traditional sense, they are weighted sums of the basis $X_{i}$, not some matrix columns.

[K.J. Barnes (1963)](https://spiral.imperial.ac.uk/bitstream/10044/1/13413/2/Barnes-KJ-1963-PhD-Thesis.pdf) seeks the spectrum differently, with the matrix projection operators.

Mysteriously, the eigenvalues will have multiplicities which can be deduced independently from the Lorentz group theory (Lorentz with "t"), without any iterations and polynomial equations. The group theory alone, however, will not get us to the eigenvector equations leading to the Lorenz gauge condition for spin 1 (Lorenz without "t").

**Exercise 3.** Verify Shankland's spectral results, esp. the case with one vector and one spinor index: "... we find, together with their antiparticles, the following groups of particles: a quadruplet, and two doublets."

Hint:

According to group theory, combining indices means taking "tensor products $(m,n)\otimes (k,l)$ of the Lorentz irreps", clf. Weinberg's QFT, Vol. 1, pages 229-233. What is relevant here is that each such an irrep constitutes a set of subspaces with multiplicities $2s+1$ for $s=|m-n|, |m-n+1|,\,\ldots, m+n$.

- $(0,0)$: A scalar. Shankland's singlet: A single subspace with eigenvalue multiplicity $2\cdot 0+1=1$.

- $(\frac{1}{2},\frac{1}{2})$: A single four-vector index. Shankland's singlet and triplet: two subspaces $0, 1$ with multiplicities 1 and 3.

- $(\frac{1}{2},0)\oplus (0,\frac{1}{2})$: A full single spinor index. Shankland's doublet and its antidoublet: $\frac{1}{2},  \frac{1}{2}$ subspaces with multiplicites 2 and 2.

- $(1,1)$: Two symmetric tensor indices. A mismatch with Shankland's pentuplet, triplet, and two singlets: Subspaces $0, 1, 2$ with multiplicities 1, 3, and 5. Where is the missing singlet? In a symmetric two-index tensor case, to remove a singlet also means to make the tensor traceless, so the group theory still matches Shankland with this assumption.

- $(1,0)\oplus (0,1)$: Two asymmetric tensor indices. Shankland's two particle triplets: Subspaces $1$ and $1$ with multiplicities 3 and 3.

- $(\frac{1}{2},\frac{1}{2}) \otimes \Big((\frac{1}{2},0)\oplus (0,\frac{1}{2})\Big)$, i.e. combining a vector and a spinor index?

The last case, spin $\frac{3}{2})$, splits into a spinor and $(1,\frac{1}{2}) \oplus (\frac{1}{2},1)$, clf. Weinberg's QFT, Vol. 1, page 232. The latter brings subspaces $\frac{1}{2}$ and $\frac{3}{2}$ with multiplicities 2 and 4, along with their "antisubspaces". All of this combined perfectly matches the result of Shankland.

Note that the construction of algebras is skipped, but it is not trivial. For spin $\frac{3}{2}$, Shankland had to spot that the combination $\gamma_{\mu}p^{\mu}$ acted independently of $p$, $\gamma$, and $g$. This has effectively doubled the basis dimension of the vector-spinor algebra from 5 to 10.

## Other Relevant Algebras

One can find some other mildly successful uses/hints of tensor algebras in [Phys. Rev. 106, 1345 (1957)](https://journals.aps.org/pr/abstract/10.1103/PhysRev.106.1345); [Nuovo Cimento, 43, 475 (1966)](https://link.springer.com/article/10.1007/BF02752873); [Nuovo Cimento 47, 145 (1967)](https://link.springer.com/article/10.1007/BF02818340); [Phys. Rev. 153, 1652 (1967)](https://journals.aps.org/pr/abstract/10.1103/PhysRev.153.1652); [Phys. Rev. 161, 1631 (1967)](https://journals.aps.org/pr/abstract/10.1103/PhysRev.161.1631); [Phys. Rev. D 8, 2650 (1973)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.8.2650); [Nuovo Cimento 28, 409 (1975)](https://inspirehep.net/literature/98459); [Phys. Lett. B 301 4 339 (1993)](https://arxiv.org/abs/hep-th/9212008); [Phys. Rev. C 64, 015203 (2001)](https://arxiv.org/abs/hep-ph/0103172); [Phys. Rev. D 64, 125013 (2001)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.64.125013); [Hadronic J. 26, 351 (2003)](https://www.imath.kiev.ua/~nikitin/PAPER26.pdf); [Phys. Rev. D 67, 085021 (2003)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.67.085021); [Phys. Rev. D 67, 125011 (2003)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.67.125011); [Nucl. Phys. B724, 453 (2005)](https://arxiv.org/abs/hep-th/0505255); [Phys. Rev. D 74, 084036 (2006)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.74.084036); [P. Cvitanović (2008)](https://birdtracks.eu/); [V. Monchiet and G. Bonnet (2010)](https://royalsocietypublishing.org/doi/10.1098/rspa.2010.0149); [Phys. Rev. D 97, 115043 (2018)](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.97.115043); [SUGRA and CDC](https://news.stonybrook.edu/facultystaff/qa-with-breakthrough-prize-winner-peter-van-nieuwenhuizen/)...

It is tough to read this literature, and the results may not always justify the complexity.

## Why Shankland?

To sum up, we are given a field with tensor/spinor indices and their permutation symmetries. A Lorentz-invariant operator is then constructed. It may serve as a quadratic form for the field, which in turn may serve later in building invariant physics. A spin content of the field is discovered as the eigenvalue multiplicities of that operator. One test of this formalism confirms that removing spin 0 from a vector field leads to "apesanteur" $A$ aka vector potential.

Considering the vast literature on group theory, irreducible representations, angular momentum, higher-spin field theories, spin projection operators, tensors, spinors, Weyl, Wigner, and Weinberg... Shankland's system is the closest thing to the assembly language of nature.

<div class="imgcontainer">
<a style="font-size: 1.5rem;" href="https://youtu.be/Y183gJQ9yCY?t=20">Sign the contract big boy...</a>
<br><br>
<a href="https://voi.id/en/memori/185165">
<img src="imgs/IronMike-min-rs.png" title="Iron Mike" width="60%">
</a>
</div>
