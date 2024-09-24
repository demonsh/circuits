pragma circom 2.0.0;
// include "../node_modules/circomlib/circuits/eddsaposeidon.circom";

/*This circuit template checks that c is the multiplication of a and b.*/  

template Multiplier2 () {  

   // Declaration of signals.  
   signal input value;  
   signal input receiver;
   signal input nonce;
   signal output nullifier;


   // Constraints.  
   component pos = Poseidon(3);
   pos.inputs[0] <== value;
   pos.inputs[1] <== receiver;
   pos.inputs[2] <== nonce;
   c <== pos.out;     
}